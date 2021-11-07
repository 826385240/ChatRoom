package roomMgr

import (
	cmd "ChatRoom/src/cmdid"
	"ChatRoom/src/lib/logic"
	"ChatRoom/src/lib/tcptask"
	"ChatRoom/src/protoout/chat"
	"ChatRoom/src/serverlogic/roleMgr"
	"container/list"
	"strings"
	"unsafe"
)

type Room struct {
	roomUuid uint64
	roomRoles map[string]bool
	allMessages *list.List
}

type roomManager struct {
	generateId uint64
	uuid2room map[uint64]string
	allRooms map[string]Room
	role2room map[string]string
}

var forbidWords=[...]string{"4r5e", "5h1t", "5hit", "a55", "anal","anus", "ar5e", "arrse", "arse", "ass","ass-fucker", "asses", "assfucker", "assfukka", "asshole","assholes", "asswhole", "a_s_s", "b!tch", "b00bs","b17ch", "b1tch", "ballbag", "balls", "ballsack","bastard", "beastial", "beastiality", "bellend", "bestial","bestiality", "bi+ch", "biatch", "bitch", "bitcher","bitchers", "bitches", "bitchin", "bitching", "bloody","blow job", "blowjob", "blowjobs", "boiolas", "bollock","bollok", "boner", "boob", "boobs", "booobs","boooobs", "booooobs", "booooooobs", "breasts", "buceta","bugger", "bum", "bunny fucker", "butt", "butthole","buttmunch", "buttplug", "c0ck", "c0cksucker", "carpet muncher","cawk", "chink", "cipa", "cl1t", "clit","clitoris", "clits", "cnut", "cock", "cock-sucker","cockface", "cockhead", "cockmunch", "cockmuncher", "cocks","cocksuck", "cocksucked", "cocksucker", "cocksucking", "cocksucks","cocksuka", "cocksukka", "cok", "cokmuncher", "coksucka","coon", "cox", "crap", "cum", "cummer","cumming", "cums", "cumshot", "cunilingus", "cunillingus","cunnilingus", "cunt", "cuntlick", "cuntlicker", "cuntlicking","cunts", "cyalis", "cyberfuc", "cyberfuck", "cyberfucked","cyberfucker", "cyberfuckers", "cyberfucking", "d1ck", "damn","dick", "dickhead", "dildo", "dildos", "dink","dinks", "dirsa", "dlck", "dog-fucker", "doggin","dogging", "donkeyribber", "doosh", "duche", "dyke","ejaculate", "ejaculated", "ejaculates", "ejaculating", "ejaculatings","ejaculation", "ejakulate", "f u c k", "f u c k e r", "f4nny","fag", "fagging", "faggitt", "faggot", "faggs","fagot", "fagots", "fags", "fanny", "fannyflaps","fannyfucker", "fanyy", "fatass", "fcuk", "fcuker","fcuking", "feck", "fecker", "felching", "fellate","fellatio", "fingerfuck", "fingerfucked", "fingerfucker", "fingerfuckers","fingerfucking", "fingerfucks", "fistfuck", "fistfucked", "fistfucker","fistfuckers", "fistfucking", "fistfuckings", "fistfucks", "flange","fook", "fooker", "fuck", "fucka", "fucked","fucker", "fuckers", "fuckhead", "fuckheads", "fuckin","fucking", "fuckings", "fuckingshitmotherfucker", "fuckme", "fucks","fuckwhit", "fuckwit", "fudge packer", "fudgepacker", "fuk","fuker", "fukker", "fukkin", "fuks", "fukwhit","fukwit", "fux", "fux0r", "f_u_c_k", "gangbang","gangbanged", "gangbangs", "gaylord", "gaysex", "goatse","God", "god-dam", "god-damned", "goddamn", "goddamned","hardcoresex", "hell", "heshe", "hoar", "hoare","hoer", "homo", "hore", "horniest", "horny","hotsex", "jack-off", "jackoff", "jap", "jerk-off","jism", "jiz", "jizm", "jizz", "kawk","knob", "knobead", "knobed", "knobend", "knobhead","knobjocky", "knobjokey", "kock", "kondum", "kondums","kum", "kummer", "kumming", "kums", "kunilingus","l3i+ch", "l3itch", "labia", "lmfao", "lust","lusting", "m0f0", "m0fo", "m45terbate", "ma5terb8","ma5terbate", "masochist", "master-bate", "masterb8", "masterbat*","masterbat3", "masterbate", "masterbation", "masterbations", "masturbate","mo-fo", "mof0", "mofo", "mothafuck", "mothafucka","mothafuckas", "mothafuckaz", "mothafucked", "mothafucker", "mothafuckers","mothafuckin", "mothafucking", "mothafuckings", "mothafucks", "mother fucker","motherfuck", "motherfucked", "motherfucker", "motherfuckers", "motherfuckin","motherfucking", "motherfuckings", "motherfuckka", "motherfucks", "muff","mutha", "muthafecker", "muthafuckker", "muther", "mutherfucker","n1gga", "n1gger", "nazi", "nigg3r", "nigg4h","nigga", "niggah", "niggas", "niggaz", "nigger","niggers", "nob", "nob jokey", "nobhead", "nobjocky","nobjokey", "numbnuts", "nutsack", "orgasim", "orgasims","orgasm", "orgasms", "p0rn", "pawn", "pecker","penis", "penisfucker", "phonesex", "phuck", "phuk","phuked", "phuking", "phukked", "phukking", "phuks","phuq", "pigfucker", "pimpis", "piss", "pissed","pisser", "pissers", "pisses", "pissflaps", "pissin","pissing", "pissoff", "poop", "porn", "porno","pornography", "pornos", "prick", "pricks", "pron","pube", "pusse", "pussi", "pussies", "pussy","pussys", "rectum", "retard", "rimjaw", "rimming","s hit", "s.o.b.", "sadist", "schlong", "screwing","scroat", "scrote", "scrotum", "semen", "sex","sh!+", "sh!t", "sh1t", "shag", "shagger","shaggin", "shagging", "shemale", "shi+", "shit","shitdick", "shite", "shited", "shitey", "shitfuck","shitfull", "shithead", "shiting", "shitings", "shits","shitted", "shitter", "shitters", "shitting", "shittings","shitty", "skank", "slut", "sluts", "smegma","smut", "snatch", "son-of-a-bitch", "spac", "spunk","s_h_i_t", "t1tt1e5", "t1tties", "teets", "teez","testical", "testicle", "tit", "titfuck", "tits","titt", "tittie5", "tittiefucker", "titties", "tittyfuck","tittywank", "titwank", "tosser", "turd", "tw4t","twat", "twathead", "twatty", "twunt", "twunter","v14gra", "v1gra", "vagina", "viagra", "vulva","w00se", "wang", "wank", "wanker", "wanky","whoar", "whore", "willies", "willy", "xrated","xxx"}

var RoomManager = &roomManager{generateId:0, uuid2room: map[uint64]string{}, allRooms: map[string]Room{}, role2room: map[string]string{}}

func (this *roomManager) CreateRoom(n string) bool {
	_,ok := this.allRooms[n]

	//创建并缓存房间数据
	if !ok {
		this.generateId++
		this.allRooms[n]=Room{roomUuid: this.generateId, roomRoles: map[string]bool{}, allMessages: list.New()}
		this.uuid2room[this.generateId]=n
	}
	return !ok
}

func (this *roomManager) JoinRoom(task *tcptask.TcpTask, role string, room string) bool {
	_,ok := this.allRooms[room]
	if !ok {
		return ok
	}

	r := this.allRooms[room]
	_,ok1 := r.roomRoles[role]
	if ok1 {
		return false
	}

	r.roomRoles[role]=true
	this.role2room[role]=room

	//返回加入房间的结果
	ret1 := &chat.MSG_JoinRoom_SC{Retcode: true}
	logic.SendMsg(task, cmd.MSG_JoinRoom_SC, unsafe.Pointer(ret1))

	//进入房间广播50条消息
	ret2 := &chat.MSG_SendMessage_SC{}
	var i int = 0
	for e := r.allMessages.Front(); e!=nil; e=e.Next() {
		i=i+1
		if i > 50 {
			break
		}
		ret2.Message = append(ret2.Message, (e.Value).(string))
	}
	logic.SendMsg(task, cmd.MSG_SendMessage_SC, unsafe.Pointer(ret2))
	return true
}

func (this *roomManager) LeaveRoom(role string) bool {
	room,ok1 := this.role2room[role]
	if !ok1 {
		return ok1
	}

	_,ok2 := this.allRooms[room]
	if !ok2 {
		return ok2
	}

	r := this.allRooms[room]
	_,ok3 := r.roomRoles[role]
	if !ok3 {
		return ok3
	}

	//从房间从删除角色信息
	delete(r.roomRoles, role)
	delete(this.role2room, role)

	if len(r.roomRoles) <= 0 {
		delete(this.allRooms, room)
	}
	return true
}

func (this *roomManager) SendMessage(role string, msg string) bool {
	room,ok1 := this.role2room[role]
	if !ok1 {
		return ok1
	}

	_,ok2 := this.allRooms[room]
	if !ok2 {
		return ok2
	}

	r := this.allRooms[room]
	_,ok3 := r.roomRoles[role]
	if !ok3 {
		return ok3
	}

	for i:=0; i<len(forbidWords);i++ {
		msg=strings.Replace(msg,forbidWords[i],strings.Repeat("*",len(forbidWords[i])),-1)
	}

	//将玩家发送的消息添加到List尾部
	r.allMessages.PushBack(role + ":" + msg)

	//将玩家发送的消息广播房间所有人
	for k,_ := range r.roomRoles {
		task := roleMgr.RoleManager.GetTcpTask(k)
		if task != nil && task.IsWChanValid() {
			ret := &chat.MSG_SendMessage_SC{}
			ret.Message = append(ret.Message, (r.allMessages.Back().Value).(string))
			logic.SendMsg(task, cmd.MSG_SendMessage_SC, unsafe.Pointer(ret))
		}
	}
	return true
}

func (this *roomManager) GetRoomIdByRole(role string) uint64 {
	room,ok1 := this.role2room[role]
	if !ok1 {
		return 0
	}

	_,ok2 := this.allRooms[room]
	if !ok2 {
		return 0
	}

	r := this.allRooms[room]
	_,ok3 := r.roomRoles[role]
	if !ok3 {
		return 0
	}

	return r.roomUuid
}

func (this *roomManager) GetPopularWord(roomId uint64) string {
	room,ok1 := this.uuid2room[roomId]
	if !ok1 {
		return ""
	}

	r,ok2 := this.allRooms[room]
	if !ok2 {
		return ""
	}

	//遍历所有消息统计单词出现频率
	countMap := map[string]uint{}
	for e := r.allMessages.Front(); e!=nil; e=e.Next() {
		nameMsg := strings.Split((e.Value).(string), ":")
		message := strings.Replace((e.Value).(string), (nameMsg[0]+":"),"",1)
		strArray := strings.Split(message, " ")
		for i:=0; i<len(strArray); i++{
			countMap[strArray[i]]++
		}
	}

	//获取出现频率最高的单词
	var maxTimes uint = 0
	var retString string
	for k,v := range countMap {
		if v > maxTimes {
			retString = k
			maxTimes = v
		}
	}

	return  retString
}

func (this *roomManager) GetAllRooms() []string  {
	//获取所有房间的名字
	ret := make([]string,0,1)
	for k,_ := range this.allRooms{
		ret=append(ret, k)
	}
	return ret
}
