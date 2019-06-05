import React from 'react';
import {
    Form,
    FormGroup,
    Label,
    Input,
    Button
} from 'reactstrap'

export default function Login (props) {
    return (
        <Form>
            <FormGroup>
                <Label for="emailField">Email</Label>
                <Input type="email" name="email" id="emailField" placeholder="Email"></Input>
            </FormGroup>
            <FormGroup>
                <Label for="passwordField">Password</Label>
                <Input type="password" name="password" id="passwordField" placeholder="Password"></Input>
            </FormGroup>
            <FormGroup check>
                <Label check>
                    <Input type="checkbox"/>{' '}
                    Remember Me
                </Label>
                <Label>Forgot Password?</Label>
            </FormGroup>
            <Button>Login</Button>
        </Form>
    )
}