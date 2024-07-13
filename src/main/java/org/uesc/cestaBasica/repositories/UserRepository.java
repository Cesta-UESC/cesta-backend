package org.uesc.cestaBasica.repositories;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;
import org.uesc.cestaBasica.entities.User;

@Repository
public interface UserRepository extends JpaRepository<User, Integer> {
}
