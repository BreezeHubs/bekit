package wx

import "fmt"

func main() {
	appid := "wx4f4bc4dec97d474b"
	sessionKey := "tiihtNczf5v6AKRyjwEUhQ=="

	encryptedData := "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZM\nQmRzooG2xrDcvSnxIMXFufNstNGTyaGS\n9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+\n3hVbJSRgv+4lGOETKUQz6OYStslQ142d\nNCuabNPGBzlooOmB231qMM85d2/fV6Ch\nevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6\n/1Xx1COxFvrc2d7UL/lmHInNlxuacJXw\nu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn\n/Hz7saL8xz+W//FRAUid1OksQaQx4CMs\n8LOddcQhULW4ucetDf96JcR3g0gfRK4P\nC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB\n6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns\n/8wR2SiRS7MNACwTyrGvt9ts8p12PKFd\nlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYV\noKlaRv85IfVunYzO0IKXsyl7JCUjCpoG\n20f0a04COwfneQAGGwd5oa+T8yO5hzuy\nDb/XcxxmK01EpqOyuxINew=="

	iv := "r7BXXKkLb8qrSNn05n0qiA=="

	bwx, err := NewBeWXTools(appid, sessionKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	s, err := bwx.DecryptWXUserInfo(encryptedData, iv)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("result:%+v", s)
	}
}
