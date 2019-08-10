package auth

var _usersMap = map[string]string{} //[]string{}

func Login(name string){
	
	_, isExists := _usersMap[name]
	
	if !isExists{
		_usersMap[name] = name
	}
}


