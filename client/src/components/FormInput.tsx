import React, { HTMLInputTypeAttribute, forwardRef } from 'react';
import { FieldError } from 'react-hook-form';
import styled from 'styled-components';

const FormInputContainer = styled.div``;

const Label = styled.label``;

const InputContainer = styled.div``;

const Input = styled.input``;

type FormInputProps = {
  label: string;
  isRequired?: boolean;
  type: HTMLInputTypeAttribute;
  error: FieldError | undefined;
};

export const FormInput = forwardRef<HTMLInputElement, FormInputProps>(
  ({ label, type, error }, ref) => {
    return (
      <FormInputContainer>
        <Label>{label}</Label>
        <InputContainer>
          <Input type={type} ref={ref} />
        </InputContainer>
        <>{error}</>
      </FormInputContainer>
    );
  }
);

FormInput.displayName = 'FormInput';
