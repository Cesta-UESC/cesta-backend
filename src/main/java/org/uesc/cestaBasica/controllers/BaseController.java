package org.uesc.cestaBasica.controllers;

import java.util.HashMap;
import java.util.Map;
import java.util.Random;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.validation.FieldError;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.uesc.cestaBasica.api.ApiException;

public class BaseController {
	Logger logger = LoggerFactory.getLogger(BaseController.class);
	
    @ResponseStatus(HttpStatus.BAD_REQUEST)
    @ExceptionHandler(MethodArgumentNotValidException.class)
    public Map<String, String> handleValidationExceptions(MethodArgumentNotValidException ex) {
        var errors = new HashMap<String, String>();
        ex.getBindingResult().getAllErrors().forEach((error) -> {
            String fieldName = ((FieldError) error).getField();
            String errorMessage = error.getDefaultMessage();
            errors.put(fieldName, errorMessage);
        });
        return errors;
    }
    
    @ResponseStatus(HttpStatus.UNPROCESSABLE_ENTITY)
    @ExceptionHandler(ApiException.class)
    public Map<String, String> handleApiExceptions(ApiException ex) {
        var errors = new HashMap<String, String>();
        errors.put("code", ex.getCode());
        errors.put("details", ex.getMessage());
        return errors;
    }
    
    @ResponseStatus(HttpStatus.INTERNAL_SERVER_ERROR)
    @ExceptionHandler(Exception.class)
    public Map<String, String> handleInternalExceptions(Exception ex) {
        var errors = new HashMap<String, String>();
        	
        var possible_values = "ABCDEFGHJKLMNPQRSTUVWXYZ0123456789";
        var qtd = 6;
        
        var code = new StringBuilder();
        
        for (int i = 0; i < qtd; i++) {
			code.append(possible_values.charAt(new Random().nextInt(0, possible_values.length())));
		}
        
        errors.put("details", "Erro inexperado, contate o suporte e informe o código " + code + ".");
        logger.error("[" + "" + code + "]", ex);
        return errors;
    }
}
