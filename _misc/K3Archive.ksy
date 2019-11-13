meta:
  id: k3_archive
  file-extension: K3Archive
  endian: le
seq:
  - id: hdr
    type: ros_chunk_file_header1
  - id: chunks_info
    type: ros_chunk_file_header2
    repeat: expr
    repeat-expr: hdr.max_chunks
enums:
  chunk:
    1:   comment
    51:  signal_spec
    52:  signal_header
    53:  signal_data
    54:  signal_balancing
    101: rms_spec
    102: rms_data
    151: center_data
types:
  ros_chunk_file_header1:
    seq:
      - id: chunk_file_id
        contents: [0x12, 0x13]
      - id: max_chunks
        type: u2
      - id: num_chunks
        type: u2
  ros_chunk_file_header2:
    seq:
      - id: chunk_id
        type: u4
        enum: chunk
      - id: chunk_ver
        type: u1
      - id: offset
        type: u4
      - id: size
        type: u4
    # встраиваем данные
    instances:
      data:
        pos: offset
        # создаём под-стрим - signal_data_ использует eos
        size: size
        type:
          switch-on: chunk_id
          cases:
            'chunk::comment':          comment(chunk_ver)
            'chunk::signal_spec':      signal_spec_v1
            'chunk::signal_header':    signal_header_v1
            'chunk::signal_data':      signal_data_v1
            'chunk::signal_balancing': signal_balancing_v1
            'chunk::rms_spec':         rms_spec_v1
            'chunk::rms_data':         rms_data_v1
            'chunk::center_data':      center_data_v1
  comment:
    params:
      - id: ver
        type: u1
    seq:
      - id: comment
        type:
          switch-on: ver
          cases:
            1: comment_v1
  comment_v1:
    seq:
      - id: text
        type: strz
        size: 26
        encoding: ASCII
  signal_spec_v1:
    seq:
      - id: reg_mode
        type: u1
      - id: reg_freq
        type: u4
      - id: lpf_freq
        type: u4
      - id: points_count
        type: u4
      - id: channels_count
        type: u1
      - id: channels_helper
        type: helper
        repeat: expr
        repeat-expr: 6  # SuperMaxRegChannelsCountPlus2
      - id: src_first_marker_index
        type: u4
      - id: data_by_fames
        type: b1
    types:
      helper:
        seq:
          - id: type
            type: u1
          - id: direction
            type: u1
  signal_header_v1:
    seq:
      - id: reg_channel_num
        type: u1
      - id: data_type
        type: u1
      - id: data_units
        type: u1
      - id: points_count
        type: u4
      - id: dx
        type: f4
      - id: coeff
        type: f4
  # [каналы][точки (+ фазы)], int16_t или float_t в зависимости от signal_spec
  signal_data_v1:
    seq:
      - id: points
        type:
          switch-on: 0  # TODO: signal_spec.reg_mode
          cases:
            0: s2  # сигнал
            1: f4  # спектр
            _: u1  # чтобы не было бесконечного цикла из-за eos
        repeat: eos
  signal_balancing_v1:
    seq:
      - id: rpm
        type: f4
      - id: harmonics
        type: harmonics
        repeat: expr
        repeat-expr: 4  # SuperMaxRegChannelsCount
      - id: load
        type: load
    types:
      harmonics:
        seq:
          - id: ampl
            type: f4
          - id: phase
            type: f4
      load:
        seq:
          - id: weight
            type: f4
          - id: angle
            type: f4
  rms_spec_v1:
    seq:
      - id: calced_params_start
        type: u1
      - id: calced_params_next
        type: u1
        repeat: expr
        repeat-expr: 7  # RMS_CALCED_PARAMS_COUNT
      - id: param_for_vibro_channels
        type: u1
        repeat: expr
        repeat-expr: 3  # MaxVibroChannelTypesCount
      - id: values_presence
        type: u2
        repeat: expr
        repeat-expr: 3  # MaxOrientationsCount
  # [точки, 14][направления, 3][рассчитываемые параметры, 7], float_t
  rms_data_v1:
    seq:
      - id: points
        type: orients
        repeat: expr
        repeat-expr: 14  # MaxRmsPointsCount
    types:
      orients:
        seq:
          - id: orients
            type: params
            repeat: expr
            repeat-expr: 3  # MaxOrientationsCount
      params:
        seq:
          - id: params
            type: f4
            repeat: expr
            repeat-expr: 7  # RMS_CALCED_PARAMS_COUNT
  center_data_v1:
    seq:
      - id: rpm_zone
        type: u1
      - id: coupling_diameter
        type: u2
      - id: distance1
        type: u4
      - id: distance2
        type: u4
      - id: osev
        type: f4
        repeat: expr
        repeat-expr: 4  # MaxCenterMsrPointsCount
      - id: poper
        type: f4
        repeat: expr
        repeat-expr: 4  # MaxCenterMsrPointsCount
      - id: missing_point
        type: u1
