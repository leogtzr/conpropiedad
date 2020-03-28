package com.conpropiedad.domain;

import lombok.*;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.mapping.MongoId;

import java.util.*;

import static java.util.Arrays.asList;
import static java.util.stream.Collectors.toUnmodifiableList;

@Data
@ToString
@NoArgsConstructor
@AllArgsConstructor
@EqualsAndHashCode
@Document
public class Word {

    @MongoId
    private String word;
    private String meaning;
    private Set<String> tags;

    public void setTags(final String tags) {
        this.tags = new TreeSet<>(asList(tags.split(",")).stream().distinct().collect(toUnmodifiableList()));
    }

    public List<String> getTags() {
        return new ArrayList<>(this.tags);
    }
}
