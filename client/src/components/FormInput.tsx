import React, { FC, HTMLInputTypeAttribute, InputHTMLAttributes } from 'react';
import styled from 'styled-components';

const FormInputContainer = styled.div``;

const Label = styled.label``;

const InputContainer = styled.div``;

const Input = styled.input``;

type FormInputProps = {
  label: string;
  isRequired?: boolean;
  value: InputHTMLAttributes<unknown>['value'];
  type: HTMLInputTypeAttribute;
};

export const FormInput: FC<FormInputProps> = ({ label, type, value }) => {
  return (
    <FormInputContainer>
      <Label>{label}</Label>
      <InputContainer>
        <Input type={type} value={value} />
      </InputContainer>
    </FormInputContainer>
  );
};
