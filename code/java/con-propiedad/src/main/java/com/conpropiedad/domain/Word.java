package com.conpropiedad.domain;

import lombok.*;
import org.springframework.data.mongodb.core.mapping.Document;

import java.util.*;

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
        this.tags = new TreeSet<>(Arrays.asList(tags.split(",")));
    }

    public List<String> getTags() {
        return new ArrayList<>(this.tags);
    }
}
