package org.uesc.cestaBasica.repositories;

import java.util.Optional;

import org.springframework.boot.autoconfigure.data.web.SpringDataWebProperties.Sort;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;
import org.uesc.cestaBasica.entities.Cesta;

@Repository
public interface CestaRepository extends JpaRepository<Cesta, Long> {
	public Optional<Cesta> findFirstByNomeLike(String nome);
}
