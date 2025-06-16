CREATE TABLE IF NOT EXISTS
    cell (
        id SERIAL PRIMARY KEY,
        lat REAL NOT NULL,
        long REAL NOT NULL,
        cell_tech VARCHAR(7),
        cell_identity bigint,
        plmn VARCHAR(6),
        lac SMALLINT CHECK (lac BETWEEN 1 AND 65534),
        rac SMALLINT CHECK (rac BETWEEN 0 AND 255),
        tac SMALLINT CHECK (tac BETWEEN 0 AND 65534),
        frequency_band VARCHAR(8),
        arfcn INTEGER,
        frequency_mhz NUMERIC(8, 3),
        -- 2G Metrics
        rxlev SMALLINT CHECK (rxlev BETWEEN 0 AND 63), -- RxLev (0-63)
        rxqual SMALLINT CHECK (rxqual BETWEEN 0 AND 7), -- RxQual (0-7)
        ec_n0 NUMERIC(5, 2), -- Ec/N0 in dB (float)
        c_i NUMERIC(5, 2), -- Carrier to Interference ratio in dB
        -- 3G Metrics
        rscp INTEGER CHECK (rscp BETWEEN -120 AND -40), -- RSCP in dBm
        -- 4G Metrics
        rsrp INTEGER CHECK (rsrp BETWEEN -140 AND -44), -- RSRP in dBm
        rsrq NUMERIC(5, 2) CHECK (rsrq BETWEEN -20 AND -3), -- RSRQ in dB
        sinr NUMERIC(5, 2), -- SINR in dB
        generated_at timestamp(0)
        with
            time zone,
            created_at timestamp(0)
        with
            time zone NOT NULL DEFAULT NOW()
    );