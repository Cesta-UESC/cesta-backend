package org.uesc.cestaBasica.api;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Size;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@AllArgsConstructor
@NoArgsConstructor
@Data
public class CestaData {
	@NotBlank
	@Size(min = 2, max = 100, message = "Nome deve tamanho entre 2 e 100 caracteres")
	private String nome;
}
