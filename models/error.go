package model


type Error struct {
	Rcode int	`json:"code"`
	Messeage string	`json:"message"`
	Status bool	`json:"status"`
}

type Erro struct {
	Messeage string	`json:"message"`
}

func Err(errIn int) Error {

	switch errIn {
	case 100:
		return Error{Rcode:100,Messeage: "wrong format",Status:false}
		break
	case 402:
		return Error{Rcode:402 ,Messeage: "user not found",Status:false}
		break

	case 200:
		return Error{Rcode:200 ,Messeage: "success",Status:true}
		break

	case 401:
		return Error{Rcode:401 ,Messeage: "token is expired",Status:false}
		break

	case 410:
		return Error{Rcode:410 ,Messeage: "wrong password",Status:false}
		break


	default:
		return Error{Rcode:400 ,Messeage: "invalid error",Status:false}
		break
	}

	return Error{Rcode:400 ,Messeage: " invalid error",Status:false}


}
