---
title: 'Faster output in Java for Competitive Programming'
date: 2020-07-06
permalink: /posts/2020/07/fast-output-java-for-competitive-programming/
tags:
  - competitive-programming
---

I was solving some Kattis problem and I kept running into a TLE(Time Limit Exceeded).
I thought my solution couldn't be optimized anymore, I had even implemented my own scanner instead of using the System.in and I used hashmaps and sets in my solution but still with no success. 
After a lot of trial and error, I found out the culprit me was using System.out.println to print the answer, I replaced that with a PrintWriter instead and all of a sudden I pass the test with 0.62 seconds. 

That drove me to run some benchmarks on using System.in vs a PrintWriter with System.out and compare the results. This current test was ran on a MacBook Pro 13-inch 2019: 

```java
public class Benchmarks {
	public static void main(String[] args) {
		Instant start = Instant.now();

		for (int i = 0; i <= 2000000; i++)
			System.out.println(i);

		Duration printStreamTime = Duration.between(start, Instant.now());

		PrintWriter writer = new PrintWriter(new BufferedWriter(new OutputStreamWriter(System.out)));

		System.gc();
		try {
			Thread.sleep(1000L);
		} catch (InterruptedException ie) {
		}

		start = Instant.now();

		for (int i = 0; i <= 2000000; i++)
			writer.println(i);

		writer.close();

		Duration printWriterTime = Duration.between(start, Instant.now());

		double percent = ((double) printWriterTime.toNanos() / printStreamTime.toNanos() * 100.0);
		System.out.println("PrintStream time = " + printStreamTime.toMillis());
		System.out.println("PrintWriter time = " + printWriterTime.toMillis());
		System.out.println(percent);
	}
}
```

PrintStream time = 6417 milliseconds,
PrintWriter time = 3366 milliseconds 

System.out.println uses a PrintStream which prints a stream of bytes, whereas a PrintWriter prints a stream of characters, that caused on average the Printwriter on average to be 52% faster than using the System.out.println!
So next time you're stuck with a TLE, maybe try using a printwriter and hopefully you'll pass the tests.
