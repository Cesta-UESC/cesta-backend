package org.uesc.cestaBasica.api.errors;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class InternalServerErrorModel {
    private String details;
}