import React, { Component } from 'react';
import { withRouter } from 'react-router-dom';
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

    const words = this.state.words.map(w => {
      return <ul><li><b><i>{w.word}:</i></b> {w.meaning}    <i>[{w.tags.join(", ")}]</i></li></ul>
    });

    return <div>
      <AppNavbar/>
      <h2>Search for insultos</h2>;
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
      {words}
      </div>
  }
}

export default withRouter(Insultos);