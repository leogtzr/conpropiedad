package com.conpropiedad.domain;

import lombok.*;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

@Data
@ToString
@NoArgsConstructor
@AllArgsConstructor
@EqualsAndHashCode
public class Word {

    private String word;
    private String meaning;
    private List<String> tags;

    public void setTags(final String tags) {
        this.tags = Arrays.asList(tags.split(","));
    }

    public List<String> getTags() {
        return new ArrayList<>(this.tags);
    }
}
