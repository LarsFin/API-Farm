using System;
using System.Globalization;
using ApiFarm.Models.Impl;
using Newtonsoft.Json;

namespace ApiFarm.Utils.Impl
{
    /// <summary>
    /// Responsible for converting the DateTime of a <see cref="VideoGame"/>.
    /// </summary>
    public class VideoGameDateTimeConverter : JsonConverter<DateTime>
    {
        private const string DateFormat = "dd/MM/yyyy";

        /// <summary>
        /// Reads the string value of the DateTime and attempts to parse it against the format 'dd/MM/yyyy'.
        /// </summary>
        /// <param name="reader">The <see cref="JsonReader"/> to read from.</param>
        /// <param name="objectType">Type of the object.</param>
        /// <param name="existingValue">The existing value of object being read.</param>
        /// <param name="hasExistingValue">Whether an existing value has been set for the object being read.</param>
        /// <param name="serializer">The calling serializer.</param>
        /// <returns>The parsed value.</returns>
        public override DateTime ReadJson(JsonReader reader, Type objectType, DateTime existingValue, bool hasExistingValue, JsonSerializer serializer)
        {
            var stringValue = reader.Value as string;

            if (string.IsNullOrEmpty(stringValue))
            {
                return default;
            }

            if (DateTime.TryParseExact(stringValue, DateFormat, CultureInfo.InvariantCulture, DateTimeStyles.None, out var dateTime))
            {
                return dateTime;
            }

            throw new NotSupportedException($"The provided date_released '{stringValue}' is invalid.");
        }

        /// <summary>
        /// Writes the <see cref="DateTime"/> as a string to a Json result.
        /// </summary>
        /// <param name="writer">The <see cref="JsonWriter"/> to write to.</param>
        /// <param name="value">The value.</param>
        /// <param name="serializer">The calling serializer.</param>
        public override void WriteJson(JsonWriter writer, DateTime value, JsonSerializer serializer)
        {
            writer.WriteValue(value.ToString(DateFormat));
        }
    }
}
