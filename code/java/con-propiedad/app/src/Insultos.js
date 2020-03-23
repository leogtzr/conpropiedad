import React, { Component } from 'react';
import { withRouter } from 'react-router-dom';
import './App.css';
import { Button, Container, Form, FormGroup, Input, Label } from 'reactstrap';
import AppNavbar from './AppNavbar';

class Insultos extends Component {

  constructor(props) {
    super(props);
    this.state = {
      value: '',
      words: []
    };
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    this.setState({value: event.target.value});
  }
  
  handleSubmit(event) {

    fetch('words/' + this.state.value)
        .then(response => response.json())
        .then(data => this.setState({words: data}));

    event.target.reset();
    event.preventDefault();
  }

  render() {
    const {words} = this.state;
    console.log(words);
    const title = <h2>Search for insultos</h2>;

    const todo = words.map(w => {
      return <ul><li>{w.word}: {w.meaning}, tags: <b>[{w.tags.join(", ")}]</b></li></ul>
    });

    return <div>
      <AppNavbar/>
      {title}
      <Container>
      <Form onSubmit={this.handleSubmit}>
        <FormGroup>
          <Label for="name">Keywords</Label>
          <Input type="text" onChange={this.handleChange} name="name" id="name"/>
          </FormGroup>
        <FormGroup>
          <Button color="primary" type="submit">Search</Button>
        </FormGroup>
      </Form>
      </Container>
      {todo}
      </div>
  }
}

export default withRouter(Insultos);