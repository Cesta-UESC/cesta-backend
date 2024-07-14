package org.uesc.cestaBasica.entities.compositeForeingKeys;

import java.io.Serializable;

import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;

@AllArgsConstructor
@NoArgsConstructor
public class SearchCityPK implements Serializable {
    protected int cityId;
    protected int searchId;
}
