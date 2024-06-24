/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   atoi.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 16:47:46 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/24 17:24:20 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func Atoi(s string) int {
	result := 0
	sign := 1
	for i := 0; i < len(s); i++ {
		if i == 0 && (s[i] == '+' || s[i] == '-') {
			if s[i] == '-' {
				sign = -1
			}
			continue
		}
		if !IsDigit(s[i]) {
			return 0
		}
		digit := int(s[i] - '0')
		result = result * 10 + digit
	}
	return sign * result
}
