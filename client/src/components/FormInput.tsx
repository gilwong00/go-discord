import { theme } from '@/theme/theme';
import React, { HTMLInputTypeAttribute, forwardRef } from 'react';
import { FieldError } from 'react-hook-form';
import styled from 'styled-components';

interface InputProps {
  backgroundColor: string;
}

interface LabelProps {
  fontSize: number;
  fontColor: string;
}

const FormInputContainer = styled.div``;

const Label = styled.label<LabelProps>`
  margin-bottom: 10px;
  font-size: ${props => props.fontSize}px;
  color: ${props => props.fontColor};
  text-transform: uppercase;
  font-weight: 700;
  line-height: 16px;
  letter-spacing: 0.24px;
`;

const InputContainer = styled.div``;

const Input = styled.input<InputProps>`
  height: 40px;
  padding: 10px;
  border: none;
  outline: none;
  width: 100%;
  border-radius: 4px;
  background-color: ${(props: InputProps) => props.backgroundColor};
`;

const Required = styled.span`
  margin-left: 5px;
  color: red;
  font-size: 8px;
`;

const ErrorText = styled.span`
  color: red;
  font-size: 14px;
  margin-left: 5px;
`;

const LabelContainer = styled.div`
  display: flex;
  flex-direction: row;
`;

type FormInputProps = {
  label: string;
  labelFontSize?: number;
  labeFontColor?: string;
  isRequired?: boolean;
  backgroundColor?: string;
  type: HTMLInputTypeAttribute;
  error: string;
};

export const FormInput = forwardRef<HTMLInputElement, FormInputProps>(
  (
    {
      label,
      labelFontSize,
      labeFontColor,
      isRequired,
      type,
      backgroundColor,
      error
    },
    ref
  ) => {
    const errorMode = error.length > 0;
    const labelColor = errorMode
      ? 'red'
      : labeFontColor ?? theme.colors.lightgray01;
    return (
      <FormInputContainer>
        <LabelContainer>
          <Label fontSize={labelFontSize ?? 14} fontColor={labelColor}>
            {label}
            {isRequired && !errorMode && <Required>*</Required>}
          </Label>
          {errorMode && <ErrorText> - {error}</ErrorText>}
        </LabelContainer>
        <InputContainer>
          <Input
            type={type}
            ref={ref}
            backgroundColor={backgroundColor ?? theme.colors.black03}
          />
        </InputContainer>
      </FormInputContainer>
    );
  }
);

FormInput.displayName = 'FormInput';
