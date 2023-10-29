package org.example.helpers;

import java.util.function.DoublePredicate;

public class Calculator<T extends Number> {
    public Double add(T x, T y) {
        return x.doubleValue() + y.doubleValue();
    }

    public Double minus(T x, T y) {
        return x.doubleValue() - y.doubleValue();
    }

    public Double mutiply(T x, T y) {
        return x.doubleValue() * y.doubleValue();
    }

    public Double divide(T x, T y) {
        if (y.doubleValue() == 0) {
            return 0.0;
        }
        return x.doubleValue() / y.doubleValue();
    }
}
