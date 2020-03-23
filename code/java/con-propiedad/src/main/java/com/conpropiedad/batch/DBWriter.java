package com.conpropiedad.batch;

import com.conpropiedad.domain.Word;
import com.conpropiedad.repository.WordRepository;
import org.springframework.batch.item.ItemWriter;
import org.springframework.beans.factory.annotation.Autowired;

import java.util.List;

public class DBWriter implements ItemWriter<Word> {

    @Autowired
    private WordRepository wordRepository;

    @Override
    public void write(final List<? extends Word> words) throws Exception {
        this.wordRepository.deleteAll().subscribe();
        this.wordRepository.saveAll(words).subscribe();
    }

}