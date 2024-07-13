package org.uesc.cestaBasica.service;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import org.uesc.cestaBasica.api.CestaData;
import org.uesc.cestaBasica.entity.Cesta;
import org.uesc.cestaBasica.repository.CestaRepository;

@Service
public class CestaService {

    @Autowired
    private CestaRepository cestaRepository;

    public List<Cesta> list() {
        return cestaRepository.findAll();
    }

    public Cesta create(CestaData cestaData) {
        return cestaRepository.save(new Cesta(cestaData.getNome()));
    }

}
