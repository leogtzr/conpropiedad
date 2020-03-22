package com.conpropiedad.batch;

import org.springframework.batch.item.ItemWriter;

import java.util.List;

// TODO: change this ... this might not be needed
public class ConsoleItemWriter<T> implements ItemWriter<T> {
    @Override
    public void write(final List<? extends T> items) {
        items.forEach(System.out::println);
    }
}