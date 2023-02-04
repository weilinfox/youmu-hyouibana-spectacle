package lib

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

var logger = logrus.WithField("log", "main")

type type155pkg byte

const (
	CLIENT_T_ACK type155pkg = iota     // 0x00
	HOST_T_ACK                         // 0x01
	INIT_ACK     type155pkg = iota + 2 // 0x04
	HOST_T                             // 0x05
	CLIENT_T                           // 0x06
	PUNCH                              // 0x07
	INIT                               // 0x08
	INIT_REQUEST                       // 0x09
	INIT_SUCCESS type155pkg = iota + 3 // 0x0b
	INIT_ERROR                         // 0x0c
	HOST_QUIT    type155pkg = iota + 5 // 0x0f
	CLIENT_QUIT                        // 0x10
	HOST_GAME    type155pkg = iota + 6 // 0x12
	CLIENT_GAME                        // 0x13
)

type data155pkg byte

const (
	GAME_SELECT         data155pkg = iota + 4 // 0x04
	GAME_INPUT          data155pkg = iota + 5 // 0x06
	GAME_REPLAY_REQUEST data155pkg = iota + 7 // 0x09
	GAME_REPLAY_MATCH                         // 0x0a
	GAME_REPLAY_DATA                          // 0x0b
	GAME_REPLAY_END                           // 0x0c
)

type match155status byte

const (
	MATCH_WAIT match155status = iota
	MATCH_ACCEPT
	// MATCH_START
	MATCH_SPECT_ACK
	MATCH_SPECT_INIT
	MATCH_SPECT_SUCCESS
	MATCH_SPECT_ERROR
)

var matchStatus = MATCH_WAIT

var th155id = [19]byte{0x57, 0x09, 0xf6, 0x67, 0xf0, 0xfd, 0x4b, 0xd0, 0xb9, 0x9a, 0x74, 0xf8, 0x38, 0x33, 0x81, 0x88, 0x00, 0x00, 0x00}
var th155SpecConf = [113]byte{0x9c, 0x00, 0x00, 0x00, 0x78, 0x9c, 0x45, 0xcd, 0x4b, 0x0a, 0x80, 0x30, 0x0c, 0x04, 0xd0,
	0x54, 0x2c, 0x8a, 0x6e, 0xba, 0xf3, 0x78, 0x12, 0xa5, 0xa0, 0xd0, 0x0f, 0xd4, 0xfa, 0x39, 0x93, 0x97, 0xd4, 0x89,
	0x88, 0xae, 0x66, 0x78, 0x49, 0x88, 0x21, 0xaa, 0x6b, 0x22, 0x9a, 0x97, 0x7e, 0xe7, 0x3c, 0x4e, 0xe8, 0x8a, 0x0c,
	0x50, 0x03, 0xed, 0x91, 0x13, 0x2b, 0x88, 0x40, 0x0b, 0x60, 0xe7, 0xe2, 0xfe, 0x2f, 0x3e, 0x5e, 0xc2, 0x03, 0x7b,
	0x2b, 0x1d, 0xf5, 0x39, 0x6e, 0x90, 0x03, 0xe7, 0xec, 0x6c, 0x1f, 0x56, 0x5f, 0x10, 0x69, 0xf5, 0x4e, 0x2a, 0xe4,
	0x66, 0xd3, 0x32, 0xc7, 0x20, 0x7c, 0x9d, 0xdd, 0xf7, 0x6d, 0x8c, 0x2e, 0x26, 0x41, 0xb9, 0x96, 0xaf, 0x37, 0x7b,
	0x90, 0x17, 0xe6}

func sockaddrIn2String(addr []byte) string {
	return fmt.Sprintf("%d.%d.%d.%d:%d", addr[2], addr[3], addr[4], addr[5], int(addr[0])<<8+int(addr[1]))
}

func littleIndia2Int(b []byte) int {
	return int(b[0]) | int(b[1])<<8 | int(b[2])<<16 | int(b[3])<<24
}

func ZlibDataDecode(l int, d []byte) string {
	if len(d) < 3 || d[0] != 0x78 || d[1] != 0x9c {
		return "NOT_ZLIB_DATA_ERROR"
	}

	b := bytes.NewBuffer(d)
	r, err := zlib.NewReader(b)
	if err != nil {
		return err.Error()
	}

	ans := make([]byte, l*2)
	n, err := r.Read(ans)
	if err != io.EOF {
		return err.Error()
	}
	r.Close()

	if l != n {
		return "ZLIB_LENGTH_NOT_MATCH_ERROR"
	}

	dataStr := ""

	i, j, s := 0, 0, 0
	for j < n {
		switch s {
		case 0:
			if ans[j] == 0x10 {
				s++
			} else {
				s = 0
			}
		case 1, 2:
			if ans[j] == 0x00 {
				s++
			} else {
				s = 0
			}
		case 3:
			if ans[j] == 0x08 {
				s++
			} else {
				s = 0
			}
		case 4:
			if i > 0 {
				dataStr += fmt.Sprint(ans[i:j-4]) + ", "
			}

			nl := littleIndia2Int(ans[j : j+4])
			i = j + 4 + nl
			dataStr += string(ans[j+4:i]) + ":"

			s = 0
		}

		j++
	}
	dataStr += fmt.Sprint(ans[i : j-4])

	return dataStr
}

func detect(buf []byte, from int, to int) {

	logger.Info(from, "->", to, " : ", buf)

	switch type155pkg(buf[0]) {

	case CLIENT_T_ACK:
		logger.Infof("CLIENT_T_ACK Client ID %d Random ID %d Time %d", buf[1], littleIndia2Int(buf[4:8]), littleIndia2Int(buf[8:12]))

	case HOST_T_ACK:
		logger.Infof("HOST_T_ACK Client ID %d Random ID %d Time %d", buf[1], littleIndia2Int(buf[4:8]), littleIndia2Int(buf[8:12]))

	case INIT_ACK:
		logger.Info("INIT_ACK")

	case HOST_T:
		logger.Infof("HOST_T Client ID %d Random ID %d Time %d", buf[1], littleIndia2Int(buf[4:8]), littleIndia2Int(buf[8:12]))

	case CLIENT_T:
		logger.Infof("CLIENT_T Client ID %d Random ID %d pad %d %d Time %d", buf[1], littleIndia2Int(buf[4:8]), littleIndia2Int(buf[8:12]), littleIndia2Int(buf[12:16]), littleIndia2Int(buf[16:20]))

	case PUNCH:
		logger.Infof("PUNCH Status %x Random ID %d IP %s", buf[1], littleIndia2Int(buf[2:6]), sockaddrIn2String(buf[6:12]))

	case INIT:
		logger.Info("INIT ", buf[1:20], " Random ID ", littleIndia2Int(buf[20:24]))

	case INIT_REQUEST:
		logger.Info("INIT_REQUEST ", buf[1:20], " Random ID ", littleIndia2Int(buf[20:24]), " Type ", buf[24:36], " Zlib ", ZlibDataDecode(littleIndia2Int(buf[36:40]), buf[40:]))

	case INIT_SUCCESS:
		logger.Info("INIT_SUCCESS ", buf[1:3], " Random ID ", littleIndia2Int(buf[4:8]), " Unknown ", buf[8:10], " IP ", sockaddrIn2String(buf[10:16]), " Pad ", buf[16:48], " Zlib ", ZlibDataDecode(littleIndia2Int(buf[48:52]), buf[52:]))
		matchStatus = MATCH_ACCEPT
		logger.Warn("==================================================")
		logger.Warn("                    MATCH ACCEPT                  ")
		logger.Warn("==================================================")

	case INIT_ERROR:
		logger.Info("INIT_ERROR Unknown ", buf[1:16], " Zlib ", ZlibDataDecode(littleIndia2Int(buf[16:20]), buf[20:]))

	case HOST_QUIT:
		logger.Info("HOST_QUIT Unknown ", buf[1:3], " Random ID ", littleIndia2Int(buf[4:8]))
		logger.Warn("==================================================")
		logger.Warn("                    HOST QUIT                     ")
		logger.Warn("==================================================")

	case CLIENT_QUIT:
		logger.Info("CLIENT_QUIT Unknown ", buf[1:3], " Random ID ", littleIndia2Int(buf[4:8]))
		logger.Warn("==================================================")
		logger.Warn("                    CLIENT QUIT                   ")
		logger.Warn("==================================================")

	case HOST_GAME:
		switch data155pkg(buf[1]) {
		case GAME_SELECT:
			logger.Info("HOST_GAME GAME_SELECT Unknown ", buf[2:5], " Match ID ", littleIndia2Int(buf[5:9]))
			/*if matchStatus < MATCH_START {
				matchStatus = MATCH_START
			}*/
		case GAME_INPUT:
			logger.Info("HOST_GAME GAME_INPUT Unknown ", buf[2], " Input ", buf[3:13], " Match ID ", littleIndia2Int(buf[13:17]), " Host frame ", littleIndia2Int(buf[17:21]), " Client frame ", littleIndia2Int(buf[21:25]))
		case GAME_REPLAY_DATA:
			start1, end1 := littleIndia2Int(buf[10:14]), littleIndia2Int(buf[14:18])
			dataEnd1 := 18 + (end1-start1)*2
			start2, end2 := littleIndia2Int(buf[dataEnd1:dataEnd1+4]), littleIndia2Int(buf[dataEnd1+4:dataEnd1+8])
			dataEnd2 := dataEnd1 + 8 + (end2-start2)*2
			logger.Info("HOST_GAME GAME_REPLAY_DATA Unknown ", buf[2:5], " Match ID ", buf[6:10], " Start ", start1, " End ", end1, " Data ", buf[18:dataEnd1], " Start ", start2, " End ", end2, " Data ", buf[dataEnd1+8:dataEnd2])
		default:
			logger.Info("HOST_GAME UNKNOWN_ID ", buf[1], " Unknown ", buf[2:])
		}

	case CLIENT_GAME:
		switch data155pkg(buf[2]) {
		case GAME_INPUT:
			logger.Info("CLIENT_GAME GAME_INPUT Client ID ", buf[1], " Unknown ", buf[3], " Input ", buf[4:14], " Match ID ", littleIndia2Int(buf[14:18]), " Host frame ", littleIndia2Int(buf[18:22]), " Client frame ", littleIndia2Int(buf[22:26]))
		case GAME_REPLAY_REQUEST:
			logger.Info("CLIENT_GAME GAME_REPLAY_REQUEST Client ID ", buf[1], " Unknown ", buf[3:6], " Match ID ", littleIndia2Int(buf[7:11]), " Unknown ", buf[11:15], " Frame ID ", buf[15:19], " Frame ID ", buf[19:23])
		default:
			logger.Info("CLIENT_GAME UNKNOWN_ID ", buf[2], " Client ID ", buf[1], " Unknown ", buf[3:])
		}

	}

}

func Sync(master, slave *net.UDPConn) {
	var slaveAddr *net.UDPAddr
	ch := make(chan int, 5)

	_, masterPortS, _ := net.SplitHostPort(master.RemoteAddr().String())
	masterPort, _ := strconv.ParseInt(masterPortS, 10, 32)

	// master -> slave
	go func() {
		defer func() {
			ch <- 1
		}()

		buf := make([]byte, 2048)

		for {
			n, err := master.Read(buf)
			if err != nil {
				logger.WithError(err).Error("master read error")
				break
			}

			if slaveAddr != nil {
				detect(buf[:n], int(masterPort), slaveAddr.Port)

				if type155pkg(buf[0]) == PUNCH { // 劫持此处所有的 punch
					buf[1] = 0x02
					buf[2], buf[3] = 0x01, 0x00
					_, err = master.Write(buf[:n])
					logger.Info("Spectator get PUNCH from host to client")
					if err != nil {
						logger.WithError(err).Error("master write punch error")
						break
					}
				} else {
					_, err = slave.WriteToUDP(buf[:n], slaveAddr)
					if err != nil {
						logger.WithError(err).Error("slave write error")
						break
					}
				}
			}
		}

	}()

	// slave -> master
	go func() {
		defer func() {
			ch <- 1
		}()

		var n int
		var err error
		buf := make([]byte, 2048)

		for {
			n, slaveAddr, err = slave.ReadFromUDP(buf)
			if err != nil {
				logger.WithError(err).Error("slave read error")
				break
			}

			detect(buf[:n], slaveAddr.Port, int(masterPort))
			_, err = master.Write(buf[:n])
			if err != nil {
				logger.WithError(err).Error("master write error")
				break
			}
		}

	}()

	var hostConn *net.UDPConn
	var matchEnd = false
	var matchId = 0
	var matchInfo []byte
	var initSuccessInfo []byte
	var initErrorInfo []byte
	var frameId = [2]int{0, 0}
	var frameRec = [2][]byte{}
	var timeId = time.Now().UnixMilli()
	var randId = rand.Int31()
	// replay request
	go func() {
		defer func() {
			ch <- 1
		}()

		hostAddr, _ := net.ResolveUDPAddr("udp", "localhost:10800")

		for hostConn == nil {
			time.Sleep(time.Millisecond * 60)

			if matchStatus != MATCH_WAIT {
				hostConn, _ = net.DialUDP("udp", nil, hostAddr)

				timeWait := 0

			bigLoop:
				for {
					time.Sleep(time.Millisecond * 33)
					timeWait += 1
					timeWait %= 30

					if timeWait == 0 {
						switch matchStatus {
						case MATCH_WAIT:
							break
						case MATCH_SPECT_ERROR:
						default:
							timeDiff := time.Now().UnixMilli() - timeId
							specData := append([]byte{byte(CLIENT_T)}, []byte{0x01, 0x71, 0x00, byte(randId), byte(randId >> 8), byte(randId >> 16), byte(randId >> 24)}...)
							specData = append(specData, []byte{0x03, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00}...)
							specData = append(specData, []byte{byte(timeDiff), byte(timeDiff >> 8), byte(timeDiff >> 16), byte(timeDiff >> 24)}...)
							logger.Info("Spectator send CLIENT_T ", specData)
							_, err := hostConn.Write(specData)
							if err != nil {
								logger.WithError(err).Error("Host conn write error")
								break
							}
						}
					} else if timeWait%2 == 0 {
						switch matchStatus {
						case MATCH_WAIT:
							break bigLoop
						case MATCH_ACCEPT:
							//	logger.Info("Spectator known match accepted")
							// case MATCH_START:
							specData := append([]byte{byte(INIT)}, th155id[:]...)
							specData = append(specData, []byte{byte(randId), byte(randId >> 8), byte(randId >> 16), byte(randId >> 24)}...)
							logger.Info("Spectator send INIT ", specData)
							_, err := hostConn.Write(specData)
							if err != nil {
								logger.WithError(err).Error("Host conn write error")
								break
							}
						case MATCH_SPECT_ACK:
							specData := append([]byte{byte(INIT_REQUEST)}, th155id[:]...)
							specData = append(specData, []byte{byte(randId), byte(randId >> 8), byte(randId >> 16), byte(randId >> 24)}...)
							specData = append(specData, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x71, 0x00, 0x71, 0x00, 0x00, 0x01}...) // 观战
							specData = append(specData, th155SpecConf[:]...)
							logger.Info("Spectator send INIT_REQUEST ", specData)
							_, err := hostConn.Write(specData)
							if err != nil {
								logger.WithError(err).Error("Host conn write error")
								break
							}
							matchStatus = MATCH_SPECT_INIT
						case MATCH_SPECT_SUCCESS:
							var specData []byte
							if matchEnd {
								specData = []byte{byte(CLIENT_GAME), 0x01, byte(GAME_REPLAY_REQUEST), 0x00, 0x00, 0x00,
									0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
									0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
							} else {
								specData = []byte{byte(CLIENT_GAME), 0x01, byte(GAME_REPLAY_REQUEST), 0x00, 0x00, 0x00,
									byte(matchId), byte(matchId >> 8), byte(matchId >> 16), byte(matchId >> 24),
									0x00, 0x00, 0x00, 0x00,
									byte(frameId[0]), byte(frameId[0] >> 8), byte(frameId[0] >> 16), byte(frameId[0] >> 24),
									byte(frameId[1]), byte(frameId[1] >> 8), byte(frameId[1] >> 16), byte(frameId[1] >> 24)}
							}
							logger.Info("Spectator send CLIENT_GAME GAME_REPLAY_REQUEST ", specData)
							_, err := hostConn.Write(specData)
							if err != nil {
								logger.WithError(err).Error("Host conn write error")
								break
							}
						}
					}
				}
			}
		}

	}()

	// replay record
	go func() {
		defer func() {
			ch <- 1
		}()

		buf := make([]byte, 2048)

		for {
			time.Sleep(time.Millisecond * 33)

			if hostConn != nil {
				n, err := hostConn.Read(buf)

				if err != nil {
					logger.WithError(err).Error("Host conn read error")
					break
				}
				switch type155pkg(buf[0]) {
				case CLIENT_T_ACK:
					logger.Info("Spectator get CLIENT_T_ACK")
				case INIT_ACK:
					logger.Info("Spectator get INIT_ACK")
					matchStatus = MATCH_SPECT_ACK
				case PUNCH:
					logger.Info("Spectator get PUNCH")
					buf[1] = 0x02
					buf[2], buf[3] = 0x01, 0x00
					_, err = hostConn.Write(buf[:n])
					if err != nil {
						logger.WithError(err).Error("Host punch reply write error")
						break
					}
				case HOST_T:
					logger.Info("Spectator get HOST_T")
					buf[0] = byte(HOST_T_ACK)
					_, err = hostConn.Write(buf[:n])
					if err != nil {
						logger.WithError(err).Error("Host host_t reply write error")
						break
					}
				case INIT_SUCCESS:
					matchStatus = MATCH_SPECT_SUCCESS
					copy(initSuccessInfo, buf[8:n])
					logger.Info("Spectator get INIT_SUCCESS ", ZlibDataDecode(littleIndia2Int(buf[48:52]), buf[52:n]))
				case INIT_ERROR:
					matchStatus = MATCH_SPECT_ERROR
					copy(initErrorInfo, buf[1:n])
					logger.Info("Spectator get INIT_ERROR ", ZlibDataDecode(littleIndia2Int(buf[16:20]), buf[20:n]))
				case HOST_GAME:
					switch data155pkg(buf[1]) {
					case GAME_REPLAY_MATCH:
						logger.Info("Spectator get HOST_GAME GAME_REPLAY_MATCH match id ", matchId, " match info ", ZlibDataDecode(littleIndia2Int(buf[17:21]), buf[21:n]))
						mid := littleIndia2Int(buf[5:9])
						if mid != matchId {
							matchId = mid
							matchEnd = false
							copy(matchInfo, buf[:n])
							frameId[0], frameId[1] = 0, 0
							frameRec[0], frameRec[1] = []byte{}, []byte{}

							logger.Warn("==================================================")
							logger.Warn("                    NEW MATCH                     ")
							logger.Warn("==================================================")
						}
					case GAME_REPLAY_DATA:
						mid := littleIndia2Int(buf[5:9])
						if mid != matchId {
							logger.Error("Spectator get invalid match id ", mid, " expect ", matchId)
						} else {
							fidS, fidE := littleIndia2Int(buf[9:13]), littleIndia2Int(buf[13:17])
							fidL := fidE - fidS
							if fidS == frameId[0] {
								frameId[0] = fidE
								frameRec[0] = append(frameRec[0], buf[17:17+fidL*2]...)
								if len(frameRec[0]) != fidE*2 {
									logger.Error("Spectator get wrong record0 length after append new data ", len(frameRec[0]), " expect ", fidE*2)
								}
							} else {
								logger.Error("Spectator get invalid start frame id ", fidS, " expect ", frameId[0])
							}
							fidS, fidE = littleIndia2Int(buf[17+fidL*2:21+fidL*2]), littleIndia2Int(buf[21+fidL*2:25+fidL*2])
							if fidS == frameId[1] {
								frameId[1] = fidE
								frameRec[1] = append(frameRec[1], buf[25+fidL*2:n]...)
								if len(frameRec[1]) != fidE*2 {
									logger.Error("Spectator get wrong record1 length after append new data ", len(frameRec[1]), " expect ", fidE*2)
								}
							} else {
								logger.Error("Spectator get invalid start frame id ", fidS, " expect ", frameId[1])
							}
							logger.Info("Spectator get HOST_GAME GAME_REPLAY_DATA match id ", matchId, " frame id ", frameId)
						}
					case GAME_REPLAY_END:
						mid := littleIndia2Int(buf[5:9])
						if mid != matchId {
							logger.Error("Spectator get invalid match id ", mid, " expect ", matchId)
						} else {
							logger.Info("Spectator get HOST_GAME GAME_REPLAY_END match id ", matchId)
							matchEnd = true
						}
					default:
						logger.Error("Spectator get invalid package ", buf[:n])
					}
				case HOST_QUIT:
					logger.Info("Spectator get HOST_QUIT")
					matchStatus = MATCH_WAIT // the end
				default:
					logger.Error("Spectator get invalid package ", buf[:n])
				}
			}
		}
	}()

	// spectator server
	go func() {
		defer func() {
			ch <- 1
		}()

		specAddr, _ := net.ResolveUDPAddr("udp", "0.0.0.0:4647")
		specConn, _ := net.ListenUDP("udp", specAddr)
		buf := make([]byte, 2048)

	serverLoop:
		for {
			n, remoteAddr, err := specConn.ReadFromUDP(buf)
			if err != nil {
				logger.WithError(err).Error("Spectator server read udp error")
				break
			}

			switch type155pkg(buf[0]) {
			case HOST_T_ACK:
				timeDiff := int64(littleIndia2Int(buf[8:12]))
				timeDiffNow := time.Now().UnixMilli() - timeId
				logger.Infof("Spectator server delay %dms", timeDiffNow-timeDiff)
			case CLIENT_T:
				repData := []byte{byte(CLIENT_T_ACK), 0x00, 0x00, 0x00}
				repData = append(repData, buf[4:8]...)
				repData = append(repData, buf[16:20]...)
				n, err = specConn.WriteToUDP(repData, remoteAddr)
				if err != nil {
					logger.WithError(err).Error("Spectator server write CLIENT_T_ACK error")
					break serverLoop
				}
				// send HOST_T
				timeDiff := time.Now().UnixMilli() - timeId
				repData = []byte{byte(HOST_T), 0x00, 0x00, 0x00, byte(randId), byte(randId >> 8), byte(randId >> 16), byte(randId >> 24),
					byte(timeDiff), byte(timeDiff >> 8), byte(timeDiff >> 16), byte(timeDiff >> 24)}
				n, err = specConn.WriteToUDP(repData, remoteAddr)
				if err != nil {
					logger.WithError(err).Error("Spectator server write HOST_T error")
					break serverLoop
				}
			case INIT:
				n, err = specConn.WriteToUDP([]byte{byte(INIT_ACK)}, remoteAddr)
				if err != nil {
					logger.WithError(err).Error("Spectator server write INIT_ACK error")
					break serverLoop
				}
			case INIT_REQUEST:
				switch buf[30] {
				case 0x6b: // 对战
					// message busy
					errData := []byte{byte(INIT_ERROR), 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x28, 0x00, 0x28, 0x00, 0x00, 0x01,
						0x1f, 0x00, 0x00, 0x00, 0x78, 0x9c, 0x13, 0x60, 0x60, 0xe0, 0x60, 0x67, 0x60, 0x60, 0xc8, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c,
						0x4f, 0x15, 0x00, 0x72, 0x59, 0x80, 0xdc, 0xa4, 0xd2, 0xe2, 0x4a, 0x46, 0x06, 0x06, 0x46, 0x00, 0x4a, 0xa5, 0x04, 0xe6}
					n, err = specConn.WriteToUDP(errData, remoteAddr)
					if err != nil {
						logger.WithError(err).Error("Spectator server write INIT_ERROR error")
						break serverLoop
					}
				case 0x71: // 观战
					var repData []byte
					switch matchStatus {
					case MATCH_SPECT_SUCCESS:
						repData = []byte{byte(INIT_SUCCESS), 0x00, 0x00, 0x00, byte(randId), byte(randId >> 8), byte(randId >> 16), byte(randId >> 24)}
						repData = append(repData, initSuccessInfo...)
					case MATCH_SPECT_ERROR:
						repData = []byte{byte(INIT_ERROR)}
						repData = append(repData, initErrorInfo...)
					default:
						// message ready
						repData = []byte{byte(INIT_ERROR), 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x29, 0x00, 0x29, 0x00, 0x00, 0x01,
							0x20, 0x00, 0x00, 0x00, 0x78, 0x9c, 0x13, 0x60, 0x60, 0xe0, 0x60, 0x67, 0x60, 0x60, 0xc8, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c,
							0x4f, 0x15, 0x00, 0x72, 0x59, 0x81, 0xdc, 0xa2, 0xd4, 0xc4, 0x94, 0x4a, 0x46, 0x06, 0x06, 0x46, 0x00, 0x51, 0x07, 0x05, 0x39}
					}
					n, err = specConn.WriteToUDP(repData, remoteAddr)
					if err != nil {
						logger.WithError(err).Error("Spectator server write INIT_SUCCESS/INIT_ERROR error")
						break serverLoop
					}
				default:
					logger.Error("Spectator server get unknown INIT_REQUEST ", buf[:n])
				}
			case CLIENT_GAME:
				switch data155pkg(buf[2]) {
				case GAME_REPLAY_REQUEST:
					mid := littleIndia2Int(buf[6:10])
					fid0, fid1 := littleIndia2Int(buf[14:18])<<1, littleIndia2Int(buf[18:22])<<1
					var repData []byte

					if mid == 0 {
						if matchId == 0 {
							break
						} else {
							// GAME_REPLAY_MATCH
							repData = matchInfo
						}
					} else if mid != matchId || matchEnd {
						// GAME_REPLAY_END
						repData = []byte{byte(HOST_GAME), byte(GAME_REPLAY_END), 0x00, 0x00, 0x00, byte(mid), byte(mid >> 8), byte(mid >> 16), byte(mid >> 24)}
					} else {
						fidE0, fidE1 := int(math.Min(float64(len(frameRec[0])), float64(fid0+48))), int(math.Min(float64(len(frameRec[1])), float64(fid1+48))) // max=24*2
						replay0, replay1 := frameRec[0][fid0+2:fidE0+2], frameRec[1][fid1+2:fidE1+2]
						fid0 >>= 1
						fid1 >>= 1
						fidE0 >>= 1
						fidE1 >>= 1
						repData = []byte{byte(HOST_GAME), byte(GAME_REPLAY_DATA), 0x02, 0x00, 0x00, byte(mid), byte(mid >> 8), byte(mid >> 16), byte(mid >> 24),
							byte(fid0), byte(fid0 >> 8), byte(fid0 >> 16), byte(fid0 >> 24), byte(fidE0), byte(fidE0 >> 8), byte(fidE0 >> 16), byte(fidE0 >> 24)}
						repData = append(repData, replay0...)
						repData = append(repData, byte(fid1), byte(fid1>>8), byte(fid1>>16), byte(fid1>>24), byte(fidE1), byte(fidE1>>8), byte(fidE1>>16), byte(fidE1>>24))
						repData = append(repData, replay1...)
					}

					n, err = specConn.WriteToUDP(repData, remoteAddr)
					if err != nil {
						logger.WithError(err).Error("Spectator server write HOST_GAME error")
						break serverLoop
					}

				default:
					logger.Error("Spectator server get unexpected CLIENT_GAME ", buf[:n])
				}
			default:
				logger.Error("Spectator server get invalid package ", buf[:n])
			}
		}
	}()

	<-ch

	logger.Info("Server terminate")
}
