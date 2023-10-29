package org.example;

import org.example.helpers.Calculator;

public class Main {
    public static void main(String[] args) {
        System.out.println("Hello and welcome!");

        Calculator<Integer> calc = new Calculator<>();
        System.out.printf("%d + %d = %.2f", 5, 7, calc.add(5, 7));
    }
}
