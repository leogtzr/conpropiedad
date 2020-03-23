import React, { Component } from 'react';
import { Link, withRouter } from 'react-router-dom';
import './App.css';
import { Button, Container, Form, FormGroup, Input, Label } from 'reactstrap';
import AppNavbar from './AppNavbar';

class Something extends Component {
  render() {
    return <p>[{this.props.valor}]</p>
  }
}

class Insultos extends Component {

  constructor(props) {
    super(props);
    this.state = {
      value: '',
      users: []
    };
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  addUser = (event) => {
    const target = event.target;
    const value = target.value;
    this.setState({
      users: [...this.state.users, <Something valor={value}/>]
    })
  }

  handleChange(event) {
    const target = event.target;
    const value = target.value;
    console.log("Hola: " + value);
  }

  handleChange(event) {
    this.setState({value: event.target.value});
  }
  
  handleSubmit(event) {
    this.setState({
        users: [...this.state.users, <Something valor={this.state.value}/>]
    });
    event.target.reset();
    event.preventDefault();
  }

  render() {
    const {item} = this.state;
    const title = <h2>Search for insultos</h2>;

    return <div>
      <AppNavbar/>
      <Container>
      <Form onSubmit={this.handleSubmit}>
        <FormGroup>
          <Label for="name">Name</Label>
          <Input type="text" onChange={this.handleChange} name="name" id="name"/>
          </FormGroup>
        <FormGroup>
          <Button color="primary" type="submit">Search</Button>
        </FormGroup>
      </Form>
      </Container>
      {this.state.users}
      </div>
  }
}

export default withRouter(Insultos);