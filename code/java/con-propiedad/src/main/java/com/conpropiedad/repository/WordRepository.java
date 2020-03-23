package com.conpropiedad.repository;

import com.conpropiedad.domain.Word;
import org.springframework.data.mongodb.repository.Query;
import org.springframework.data.mongodb.repository.ReactiveMongoRepository;
import org.springframework.stereotype.Repository;
import reactor.core.publisher.Flux;
// import reactor.core.publisher.Mono;

@Repository
public interface WordRepository extends ReactiveMongoRepository<Word, String> {
    @Query("{'tags': ?0}")
    Flux<Word> wordsWith(String tag);
}
