package org.uesc.cestaBasica.entities.compositeForeingKeys;

import java.io.Serializable;

import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;

@AllArgsConstructor
@NoArgsConstructor
public class SearchResultProductsPK implements Serializable {
    protected int productId;
    protected int cityId;
    protected int searchId;
}
