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
}
