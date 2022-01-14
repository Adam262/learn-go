package kata

import "fmt"

func Main() {
	fmt.Printf("The decimal 100 has %v x 1 bits when converted to binary form\n", countBits(100))

	fmt.Println(Is_valid_ip_easy_way("1.1.1.1"))
	fmt.Println(Is_valid_ip_easy_way("1.1.1.255"))
	fmt.Println(Is_valid_ip_easy_way("1.1.1.256"))
	fmt.Println(Is_valid_ip_easy_way("1.1.1.255.1"))
	fmt.Println(Is_valid_ip_easy_way("1.a.1.255"))

	fmt.Println(Is_valid_ip_hard_way("1.1.1.1"))
	fmt.Println(Is_valid_ip_hard_way("1.1.1.255"))
	fmt.Println(Is_valid_ip_hard_way("1.1.1.256"))
	fmt.Println(Is_valid_ip_hard_way("1.1.1.255.1"))
	fmt.Println(Is_valid_ip_hard_way("1.a.1.255"))

	fmt.Println(EncryptThis(""))
	fmt.Println(EncryptThis("A wise old owl lived in an oak"))
	fmt.Println(EncryptThis("The more he saw the less he spoke"))
	fmt.Println(EncryptThis("The less he spoke the more he heard"))
	fmt.Println(EncryptThis("Why can we not all be like that wise old bird"))
	fmt.Println(EncryptThis("Thank you Piotr for all your help"))

	fmt.Println(ListSquared(1, 250))
	fmt.Println(ListSquared(250, 500))
	fmt.Println(ListSquared(300, 600))

	fmt.Println(ValidParentheses("()()()"))
	fmt.Println(ValidParentheses("))(("))
	fmt.Println(ValidParentheses("("))
	fmt.Println(ValidParentheses(")"))
}
