package chunks

const (
	// K2
	// для раскодирования опций замера (FHOpts), формат в байте 4.2.2
	// аналог одноименных констант из файла FSUnit.h

	OptsTypeShift  = 4 // смещение в байте
	OptsTypeSkz    = 0 // СКЗ
	OptsTypeSignal = 1 // сигнал
	OptsTypePower  = 2 // ваттметрграмма

	// для сигналов
	OptsSignalTypeShift  = 2    // смещение в байте
	OptsSignalTypeMask   = 0x0C // маска
	OptsSignalTypeSignal = 0    // сигнал
	OptsSignalTypeSpectr = 1    // спектр

	OptsSignalEdIzmShift = 0    // смещение в байте
	OptsSignalEdIzmMask  = 0x03 // маска
	OptsSignalEdIzmAcc   = 0    // ускорение
	OptsSignalEdIzmVel   = 1    // скорость
	OptsSignalEdIzmDis   = 2    // перемещение
	OptsSignalEdIzmMark  = 3    // отметчик

	SuperMaxRegChannelsCountPlus2 = 6
)

const (
	// K3
	// Сигнал
	DataType_Signal uint8 = iota
	// Спектр
	DataType_Spectrum
	// Значение
	DataType_Value
)

const (
	// Виброускорение
	DataUnits_VibroAcceleration uint8 = iota
	// Виброскорость
	DataUnits_VibroVelocity
	// Виброперемещение
	DataUnits_VibroDisplacement
	// Отметчик
	DataUnits_Marker
	// Мощность полная
	DataUnits_Power_Full
	// Мощность активная
	DataUnits_Power_Active
	// Температура
	DataUnits_Temperature
	// Аудио, акустика
	DataUnits_Audio
)
