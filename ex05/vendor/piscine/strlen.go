/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   strlen.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 15:50:44 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/24 15:55:14 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

func StrLen(s string) int {
	count := 0;
	for range s {
		count++
	}
	return count
}