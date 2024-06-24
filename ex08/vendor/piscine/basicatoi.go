/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   basicatoi.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 16:18:00 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/24 16:36:03 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func BasicAtoi(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		digit := int(s[i] - '0')
		result = result * 10 + digit
	}
	return result
}