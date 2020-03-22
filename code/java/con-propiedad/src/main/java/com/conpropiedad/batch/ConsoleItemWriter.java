package com.conpropiedad.batch;

import org.springframework.batch.item.ItemWriter;

import java.util.List;

public class ConsoleItemWriter<T> implements ItemWriter<T> {
    @Override
    public void write(final List<? extends T> items) {
        items.forEach(System.out::println);
    }
}