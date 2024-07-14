package org.uesc.cestaBasica.entities.compositeForeingKeys;

import java.io.Serializable;

import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;

@AllArgsConstructor
@NoArgsConstructor
public class MinimalRationPK implements Serializable {
    protected int delimiterId;
    protected int productId;
    protected int measureId;
}
