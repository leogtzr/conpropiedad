package com.conpropiedad.web;

import com.conpropiedad.domain.Word;
import com.conpropiedad.repository.WordRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;
import reactor.core.publisher.Flux;

@RestController
public class WordController {

    @Autowired
    private WordRepository wordRepository;

    @GetMapping("/words")
    @ResponseBody
    public Flux<Word> words() {
        return this.wordRepository.findAll();
    }

    @GetMapping("/words/{tag}")
    @ResponseBody
    public Flux<Word> words(@PathVariable String tag) {
        return wordRepository.wordsWith(tag.strip().toLowerCase());
    }

}
