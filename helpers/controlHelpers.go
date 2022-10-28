package helpers

type Help struct{

}

func (h *Help)DelChar(s []rune, index int) []rune {
    return append(s[0:index], s[index+1:]...)
}