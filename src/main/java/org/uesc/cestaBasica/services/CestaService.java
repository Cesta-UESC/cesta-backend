package org.uesc.cestaBasica.services;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Sort;
import org.springframework.stereotype.Service;
import org.uesc.cestaBasica.api.ApiException;
import org.uesc.cestaBasica.api.CestaData;
import org.uesc.cestaBasica.entities.Cesta;
import org.uesc.cestaBasica.repositories.CestaRepository;

@Service
public class CestaService {

    @Autowired
    private CestaRepository cestaRepository;

    public List<Cesta> list() {
        return cestaRepository.findAll();
    }

    public Cesta create(CestaData cestaData) {
    	if (cestaRepository.findFirstByNomeLike(cestaData.getNome()).isPresent()) {
    		throw new ApiException("CestaNomeAlreadyExists", "nome já existe");
    	}
    	
    	var cesta = Cesta.builder()
    			.nome(cestaData.getNome())
    			.build();
        
    	cestaRepository.save(cesta);
    	
    	return cesta;
    }

}
