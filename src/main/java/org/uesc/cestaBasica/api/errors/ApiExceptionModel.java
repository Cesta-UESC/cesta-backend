package org.uesc.cestaBasica.api.errors;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class ApiExceptionModel {
    private String code;
    private String details;
}