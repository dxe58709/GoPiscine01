/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printstr.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 15:39:06 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/24 15:48:58 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func PrintStr(s string) {
	for _, char := range s {
		ft.PrintRune(char)
	}
}
