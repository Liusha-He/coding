package org.example;

import org.example.helpers.Calculator;

import java.util.Map;
import static java.util.Map.entry;

public class Main {
    public static void main(String[] args) {
        long startTime = System.currentTimeMillis();

        Map<Integer, Integer> map = Map.ofEntries(
                entry(5, 7),
                entry(8, 10),
                entry(13, 3),
                entry(10, 14)
        );

        Calculator<Integer> calc = new Calculator<>();
        for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
            Integer x = entry.getKey();
            Integer y = entry.getValue();

            System.out.printf("%d + %d = %.2f\n", x, y, calc.add(x, y));
            System.out.printf("%d - %d = %.2f\n", x, y, calc.minus(x, y));
            System.out.printf("%d x %d = %.2f\n", x, y, calc.mutiply(x, y));
            System.out.printf("%d / %d = %.2f\n", x, y, calc.divide(x, y));
        }

        long endTime = System.currentTimeMillis();
        long executionTime = endTime - startTime;

        System.out.printf("the code execution time = %d nanoseconds.", executionTime);
    }
}
