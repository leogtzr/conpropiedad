package com.conpropiedad.domain;

import lombok.*;
import org.springframework.data.mongodb.core.mapping.Document;

import java.util.*;

import static java.util.Arrays.stream;
import static java.util.stream.Collectors.toList;

@Data
@ToString
@NoArgsConstructor
@AllArgsConstructor
@EqualsAndHashCode
@Document
public class Word {

    private String word;
    private String meaning;
    private Set<String> tags;

    public void setTags(final String tags) {
        this.tags = new TreeSet<>(stream(tags.split(",")).map(String::toLowerCase).collect(toList()));
    }

    public List<String> getTags() {
        return new ArrayList<>(this.tags);
    }
}
