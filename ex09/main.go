/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 16:36:23 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/24 16:36:50 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.BasicAtoi2("12345"))
	fmt.Println(piscine.BasicAtoi2("0000000012345"))
	fmt.Println(piscine.BasicAtoi2("012 345"))
	fmt.Println(piscine.BasicAtoi2("Hello World!"))
}
