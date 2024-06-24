/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 15:33:25 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/24 15:34:04 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"piscine"
)
func main() {
	a := 13
	b := 2
	piscine.UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}