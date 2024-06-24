/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   basicatoi2.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 16:37:23 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/24 16:46:42 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func BasicAtoi2(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if !IsDigit(s[i]) {
			return 0
		}
		digit := int(s[i] - '0')
		result = result * 10 + digit
	}
	return result
}
