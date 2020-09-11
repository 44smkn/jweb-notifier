package vo

type TalentName string

const (
	KaitoTakahashi TalentName = "高橋海人"
	NoeruKawashima TalentName = "川島如恵留"
)

// リフレクションを利用したほうが良いかも
func TalentNameValueOf(name string) TalentName {
	switch name {
	case "高橋海人":
		return KaitoTakahashi
	case "川島如恵留":
		return NoeruKawashima
	}
	return ""
}
