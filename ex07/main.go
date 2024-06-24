/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 16:03:39 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/24 16:04:06 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "Hello World!"
	s = piscine.StrRev(s)
	fmt.Println(s)
}
