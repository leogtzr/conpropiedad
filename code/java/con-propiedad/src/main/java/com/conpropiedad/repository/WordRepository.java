package com.conpropiedad.repository;

import com.conpropiedad.domain.Word;
import org.springframework.data.mongodb.repository.ReactiveMongoRepository;
import org.springframework.stereotype.Repository;
// import reactor.core.publisher.Mono;

@Repository
public interface WordRepository extends ReactiveMongoRepository<Word, String> {
    // TODO: Mono<Word> findByTitle(String title);
}
