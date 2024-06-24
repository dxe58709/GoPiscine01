/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   strrev.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 15:59:42 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/24 16:14:51 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func StrRev(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
