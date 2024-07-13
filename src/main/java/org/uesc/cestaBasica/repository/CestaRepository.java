package org.uesc.cestaBasica.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import org.uesc.cestaBasica.entity.Cesta;

@Repository
public interface CestaRepository extends JpaRepository<Cesta, Long> {
}
