package org.uesc.cestaBasica.api;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class ApiException extends RuntimeException {
	private String code;
	private String message;
}
