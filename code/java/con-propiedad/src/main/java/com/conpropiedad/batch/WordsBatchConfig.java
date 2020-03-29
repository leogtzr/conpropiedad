package com.conpropiedad.batch;

import com.conpropiedad.domain.Word;
import org.springframework.batch.core.Job;
import org.springframework.batch.core.Step;
import org.springframework.batch.core.configuration.annotation.JobBuilderFactory;
import org.springframework.batch.core.configuration.annotation.StepBuilderFactory;
import org.springframework.batch.item.ItemProcessor;
import org.springframework.batch.item.ItemReader;
import org.springframework.batch.item.ItemWriter;
import org.springframework.batch.item.file.FlatFileItemReader;
import org.springframework.batch.item.file.LineMapper;
import org.springframework.batch.item.file.mapping.BeanWrapperFieldSetMapper;
import org.springframework.batch.item.file.mapping.DefaultLineMapper;
import org.springframework.batch.item.file.mapping.FieldSetMapper;
import org.springframework.batch.item.file.transform.DelimitedLineTokenizer;
import org.springframework.batch.item.file.transform.LineTokenizer;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.core.io.FileSystemResource;

@Configuration
public class WordsBatchConfig {

    @Autowired
    private JobBuilderFactory jobBuilderFactory;

    @Autowired
    private StepBuilderFactory stepBuilderFactory;

    @Bean
    public Job importWords(final Step step1) {
        return jobBuilderFactory.get("importWordsJob").
                start(step())
                .build();
    }

    @Bean
    public String inputFilePath(@Value("${input}") String input) {
        return input;
    }

    @Value("${input}")
    private String input;

    @Bean
    public ItemReader<Word> reader() {
        final FlatFileItemReader<Word> reader = new FlatFileItemReader<>();
        reader.setResource(new FileSystemResource(this.input));


        final LineMapper<Word> studentLineMapper = createStudentLineMapper();
        reader.setLineMapper(studentLineMapper);

        return reader;
    }

    private LineMapper<Word> createStudentLineMapper() {
        final DefaultLineMapper<Word> studentLineMapper = new DefaultLineMapper<>();

        final LineTokenizer studentLineTokenizer = createStudentLineTokenizer();
        studentLineMapper.setLineTokenizer(studentLineTokenizer);

        final FieldSetMapper<Word> studentInformationMapper = createStudentInformationMapper();
        studentLineMapper.setFieldSetMapper(studentInformationMapper);

        return studentLineMapper;
    }

    private FieldSetMapper<Word> createStudentInformationMapper() {
        final BeanWrapperFieldSetMapper<Word> studentInformationMapper = new BeanWrapperFieldSetMapper<>();
        studentInformationMapper.setTargetType(Word.class);
        return studentInformationMapper;
    }

    private LineTokenizer createStudentLineTokenizer() {
        final DelimitedLineTokenizer studentLineTokenizer = new DelimitedLineTokenizer();
        studentLineTokenizer.setDelimiter(";");
        studentLineTokenizer.setNames(new String[]{"word", "meaning", "tags"});
        return studentLineTokenizer;
    }

    @Bean
    public ItemWriter<Word> writer() {
        return new DBWriter();
    }

    @Bean
    public ItemProcessor<Word, Word> processor() {
        return item -> {
            System.out.printf("Tags have: [%d] in [%s]\n", item.getTags().size(), item.toString());
            return item;
        };
    }

    @Bean
    public Step step() {
        return stepBuilderFactory
                .get("step")
                .<Word, Word>chunk(10)
                .reader(reader())
                //.processor(processor())
                .writer(writer())
                .build();
    }

}
