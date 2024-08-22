package main

import (
	"fmt"
	"net/http"
	"time"
)

var link = "https://extra-valley.com"
var minInterval time.Duration = 10

func main() {
	for {
		resp, err := http.Get(link)
		if err != nil {
			fmt.Println("retry block ...1")
			for {
				resp, err = http.Get(link)
				if err != nil {
					for {
						fmt.Println("retry block 2")
						resp, err = http.Get(link)
						if err != nil {
							for {
								resp, err = http.Get(link)
								if err != nil {
									for {
										fmt.Println("retry block ...2")
										resp, err = http.Get(link)
										if err != nil {
											for {
												fmt.Println("retry block ...3")
												resp, err = http.Get(link)
												if err != nil {
													for {
														fmt.Println("retry block ...4")
														resp, err = http.Get(link)
														if err != nil {
															for {
																fmt.Println("retry block ...5")
																resp, err = http.Get(link)
																if err != nil {
																	for {
																		fmt.Println("retry block ...6")
																		resp, err = http.Get(link)
																		if err != nil {
																			for {
																				fmt.Println("retry block ...1")
																				resp, _ = http.Get(link)
																				fmt.Println("response header :", resp.Header)
																				time.Sleep(minInterval * time.Minute)
																			}

																		}
																		fmt.Println("response header :", resp.Header)
																		time.Sleep(minInterval * time.Minute)

																	}

																}
																fmt.Println("response header :", resp.Header)
																time.Sleep(minInterval * time.Minute)

															}

														}
														fmt.Println("response header :", resp.Header)
														time.Sleep(minInterval * time.Minute)

													}

												}
												fmt.Println("response header :", resp.Header)
												time.Sleep(minInterval * time.Minute)

											}

										}
										fmt.Println("response header :", resp.Header)
										time.Sleep(minInterval * time.Minute)

									}
								}
								fmt.Println("response header :", resp.Header)
								time.Sleep(minInterval * time.Minute)

							}
						}
						fmt.Println("response header :", resp.Header)
						time.Sleep(minInterval * time.Minute)

					}
				}
				fmt.Println("response header :", resp.Header)
				time.Sleep(minInterval * time.Minute)

			}

		}
		fmt.Println("response header :", resp.Header)
		time.Sleep(minInterval * time.Minute)
	}
}
