/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 16:18:27 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/24 16:18:47 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.BasicAtoi("12345"))
	fmt.Println(piscine.BasicAtoi("0000000012345"))
	fmt.Println(piscine.BasicAtoi("000000"))
}
